package valueobjects

import "gopkg.in/go-playground/validator.v9"

type Username struct {
	value string
}

func NewUsername(name string) (Username, error) {
	validate := validator.New()
	// Los errores de la libreria de validación pueden usarse
	// desde el momento que añado la libreria al dominio
	err := validate.Var(name, "min=5,max=64")
	if err != nil {
		return Username{}, err
	}
	return Username{name}, nil
}

func (un Username) validate() error {
	return nil
}

func (un Username) isEqualTo(username Username) bool {
	return un.value == username.Value()
}

func (un Username) Value() string {
	return un.value
}
