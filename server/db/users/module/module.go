package module

type UsersModule struct{}

func (m *UsersModule) Namespace() string {
	return "app:users"
}

func (m *UsersModule) NamespaceSeq() string {
	return "app:seq:users"
}

func (m *UsersModule) NamespaceSchema() string {
	return "app:schema:users"
}
