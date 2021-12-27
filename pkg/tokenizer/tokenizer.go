package tokenizer

// Table setup
const MAX_COLUMNS int = 130
const MAX_BUFFER int = 200
const MAX_ROWS int = 200

// State machine
const WORD_STATE int = 0
const NUMBER_STATE int = 10
const PUNC_STATE int = 20
const SPACE_STATE int = 30

// Punctuation
const PUNC string = ",.?!'[]{}/"

var Table [MAX_ROWS][MAX_COLUMNS]int

type Tokenizer struct {
	buffer []rune
	pos    int
}

func NewTokenizer() *Tokenizer {
	p := Tokenizer{pos: 0}
	if Table[MAX_ROWS-1][1] != -1 {
		p.MakeTable(&Table)
	}
	return &p
}

func (s *Tokenizer) Buffer() string {
	return string(s.buffer)
}

func (s *Tokenizer) MakeTable(table *[MAX_ROWS][MAX_COLUMNS]int) {
	InitTable(table)

	// Alpha pattern
	MakeFail(table, WORD_STATE)
	MakeSuccess(table, WORD_STATE+1)
	MarkCells(WORD_STATE, table, 'A', 'z', WORD_STATE+1)
	MarkCells(WORD_STATE+1, table, 'A', 'z', WORD_STATE+1)

	// Number pattern
	MakeFail(table, NUMBER_STATE)
	MakeSuccess(table, NUMBER_STATE+1)
	MarkCells(NUMBER_STATE, table, '0', '9', NUMBER_STATE+1)
	MarkCells(NUMBER_STATE+1, table, '0', '9', NUMBER_STATE+1)
	MakeFail(table, NUMBER_STATE+2)
	MarkCell(NUMBER_STATE, table, '.', NUMBER_STATE+2)
	MarkCell(NUMBER_STATE+1, table, '.', NUMBER_STATE+2)
	MakeSuccess(table, NUMBER_STATE+3)
	MarkCells(NUMBER_STATE+2, table, '0', '9', NUMBER_STATE+3)
	MarkCells(NUMBER_STATE+3, table, '0', '9', NUMBER_STATE+3)

	// Punctuation pattern
	MakeFail(table, PUNC_STATE)
	MakeSuccess(table, PUNC_STATE+1)
	MarkChars(PUNC_STATE, table, PUNC, PUNC_STATE+1)

	// Space pattern
	MakeFail(table, SPACE_STATE)
	MakeSuccess(table, SPACE_STATE+1)
	MarkCell(SPACE_STATE, table, ' ', SPACE_STATE+1)
}

func (s *Tokenizer) Done() bool {
	return s.pos >= len(s.buffer)
}

func (s *Tokenizer) More() bool {
	return s.pos < len(s.buffer)
}

func (s *Tokenizer) SetString(str string) {
	s.buffer = []rune(str)
	s.pos = 0
}

func (s *Tokenizer) GetToken(start_state int, token *string) bool {
	// Initialize variables
	last_success := -1
	current_pos := s.pos
	r := s.buffer[current_pos]
	if r > 127 { // Ensure character is valid
		return false
	}
	current_state := Table[start_state][r+1]

	if current_state == -1 {
		return false // When first character failed, return false immediately
	}

	// When first character is success, update last_success variable
	if IsSuccess(&Table, current_state) {
		last_success = current_pos
	}

	current_pos++ // Move to next character

	for current_pos < len(s.buffer) {
		r = s.buffer[current_pos]
		if r > 127 { // Ensure character is valid
			break
		}
		// Update current state
		current_state = Table[current_state][r+1]
		if current_state == -1 {
			break // Nowhere to go, break out and return
		}

		if IsSuccess(&Table, current_state) {
			last_success = current_pos // Update last_success
		}

		current_pos++ // Move to next character
	}

	if last_success == -1 {
		return false // No matched pattern, found illegal character
	}

	if current_state == -1 && last_success == -1 {
		return false // No matched pattern, return false
	}

	*token = string(s.buffer[s.pos : last_success+1])
	s.pos = last_success + 1
	return true
}

func (s *Tokenizer) Tokenize() *Token {
	if s.Done() {
		panic("Reach string end")
	}

	var t *Token
	tok := ""
	if s.GetToken(WORD_STATE, &tok) {
		t = NewToken(tok, 1)
	} else if s.GetToken(NUMBER_STATE, &tok) {
		t = NewToken(tok, 2)
	} else if s.GetToken(PUNC_STATE, &tok) {
		t = NewToken(tok, 3)
	} else if s.GetToken(SPACE_STATE, &tok) {
		t = NewToken(tok, 4)
	} else {
		r := s.buffer[s.pos]
		tok = string(r)
		s.pos++
		t = NewToken(tok, 0)
	}

	return t
}
