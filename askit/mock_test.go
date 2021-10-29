package askit

type (
	AskMock struct {
		InPrompt  string
		OutString string
		OutError  error
	}

	AskSecretMock struct {
		InPrompt  string
		OutString string
		OutError  error
	}

	MockAsker struct {
		AskIndex int
		AskMocks []AskMock

		AskSecretIndex int
		AskSecretMocks []AskSecretMock
	}
)

func (m *MockAsker) Output(message string) {
	// no-op
}

func (m *MockAsker) Ask(prompt string) (string, error) {
	i := m.AskIndex
	m.AskIndex++
	m.AskMocks[i].InPrompt = prompt
	return m.AskMocks[i].OutString, m.AskMocks[i].OutError
}

func (m *MockAsker) AskSecret(prompt string) (string, error) {
	i := m.AskSecretIndex
	m.AskSecretIndex++
	m.AskSecretMocks[i].InPrompt = prompt
	return m.AskSecretMocks[i].OutString, m.AskSecretMocks[i].OutError
}
