package validationkit

import "testing"

// Test for minimum numbers of characters
func TestCheckUsernameSyntaxMinimunCharacterLength(t *testing.T) {
	result := CheckUsernameSyntax("")

	if result != false {
		t.Error("Failes the minimun character check.")
	}
}

// Test for maximum number of charactes
func TestCheckUsernameSyntaxMaximumCharactersLecgth(t *testing.T) {
	result := CheckUsernameSyntax("aaaaaaaaaaaaaaaa")

	if result != false {
		t.Error("Failed the maximum character check.")
	}
}

// Test for the special characters
func TestCheckUsernameSyntaxSymbols(t *testing.T) {
	usernames := [...]string{"da!aasdad", "sdads#ad", "$qqwdqwd"}

	for _, u := range usernames {
		result := CheckUsernameSyntax(u)

		if result != false {
			t.Errorf("Failed the special charactes check. The username %s allowed", u)
		}
	}
}

// Test for the usderscore permissibility
func TestCheckUsernameSyntaxUnderscore(t *testing.T) {
	result := CheckUsernameSyntax("the_gopher")
	if result != true {
		t.Error("Failed the check to allow underscore characters.")
	}
}

// Test for the @ sign in another place besides the start of the string
func TestCheckUsernameSyntaxAtSignInsideUsername(t *testing.T) {
	result := CheckUsernameSyntax("the@gopher")
	if result != false {
		t.Error("Failed the @ sign check. The @ sign was found in another place besides the start of the string")
	}
}

// Test for 10000 random valid usernames
func TestCheckUsernameSyntaxRandomUsername(t *testing.T) {
	for i := 0; i < 10000; i++ {
		u := GenerateRandomUsername()
		r := CheckUsernameSyntax(u)
		if r != true {
			t.Error("The random username", u, "failed to pass the username check.")
			t.Fatal("Quitting!")
		}
	}
}
