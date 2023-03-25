package jsonparser

type DataWriter interface {
	GetHeaders() []string
	GetRow() Row
	Next() bool
}

func NewSingle() Single {
	return Single{
		headers: make([]string, 0),
		data:    make(Row, 0),
	}
}

type Row map[string]interface{}

type Single struct {
	headers []string
	data    Row
	counter int
}

func (s Single) GetHeaders() []string {
	return s.headers
}

func (s *Single) GetRow() Row {
	s.counter += 1
	return s.data
}

func (s Single) Next() bool {
	return s.counter < 1
}

func NewMultiple() Multiple {
	return Multiple{
		Headers: make([]string, 0),
		Rows:    make([]Row, 0),
		size:    0,
		counter: 0,
	}
}

type Multiple struct {
	Headers []string
	Rows    []Row
	counter int
	size    int
}

func (m Multiple) GetHeaders() []string {
	return m.Headers
}

func (m *Multiple) GetRow() Row {
	i := m.counter
	m.counter += 1
	return m.Rows[i]
}

func (m *Multiple) Next() bool {
	return m.counter < m.size
}

func extractHeaders(data Row) []string {
	var headers []string = make([]string, 0)
	for h := range data {
		headers = append(headers, h)
	}
	return headers
}
