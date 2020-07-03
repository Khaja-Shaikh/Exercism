package twofer

// ShareWith return the String one for name, one for me or one for you,one for me.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."

	}

	return "One for " + name + ", one for me."
}
