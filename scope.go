package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// TODO -- we're still not properly handling scope with multiple repositories
// multiple repositories are separated with a comma
// we are handling scope with one repository (one image), and push and pull permissions -- but we are just returning exactly what is asked for
func ParseScope(input string) ([]ScopeAccess, error) {

	if !strings.Contains(input, "repository") {
		return nil, errors.New("scope input string does not have a repository specified -- unable to parse what access is requested")
	}
	fmt.Println("Raw input to parse scope: " + input)
	repos := strings.Split(input, "repository:")

	accesses := make([]ScopeAccess, 0)
	for i, repo := range repos {
		if i == 0 { //skip the first entry -- it should be an empty string
			continue
		}
		// now parse the repo portion
		fmt.Println("repo: ", repo)

		allowed := make([]string, 0)
		name := ""

		// check if the ":" character in the string is immediately followed by a port number (int) before a "/"
		// if it's not, it's probably just the ":" AFTER the repo is specified
		if strings.Contains(repo, ":") {
			firstColon := strings.Index(repo, ":") //0-indexed position of :

			if len(repo) >= firstColon+1 {
				nextChar := repo[firstColon+1]
				fmt.Println(nextChar)
				restOfString := repo[firstColon+1:]

				// if there's a port number after the :
				// search for another ":", further in the repo string
				if isByteADigit(nextChar) {
					fmt.Println("byte was a digit -- restOfString: ", restOfString)

					if strings.Contains(string(restOfString), ":") {
						// get push and pull from after that
						secondColon := strings.Index(restOfString, ":")
						name = repo[0 : firstColon+secondColon+1]                           //got to the index which is the sum of the length to the first colon in whole string, and second colon in truncated string, plus 1 b/c of overlap -- todo there's probably a better way to do this in "strings" package (like find 2nd occurrence of ":") -- maybe findLastIndex, but that assumes there are no more ":")
						fmt.Println("byte was a digit -- name of reposity --- name:", name) //debugging
						endOfString := restOfString[secondColon+1:]
						fmt.Println("byte was a digit -- next part contains : -- endOfString: ", endOfString)
						if strings.Contains(endOfString, "push") { // todo this is where we would disallow certain scopes (based on "sub" we would be holding onto field of JWT -- mapping to some database with permissions per repository), if we wanted to do that
							allowed = append(allowed, "push")
						}
						if strings.Contains(endOfString, "pull") { // todo this is where we would disallow certain scopes (based on "sub" we would be holding onto field of JWT -- mapping to some database with permissions per repository), if we wanted to do that
							allowed = append(allowed, "pull")
						}
					}
				} else {
					name = repo[0:firstColon]
					fmt.Println("name: ", name)
					if strings.Contains(restOfString, "push") { // todo this is where we would disallow certain scopes (based on "sub" we would be holding onto field of JWT -- mapping to some database with permissions per repository), if we wanted to do that
						allowed = append(allowed, "push")
					}
					if strings.Contains(restOfString, "pull") { // todo this is where we would disallow certain scopes (based on "sub" we would be holding onto field of JWT -- mapping to some database with permissions per repository), if we wanted to do that
						allowed = append(allowed, "pull")
					}
				}
			} else {
				return nil, errors.New("No text after colon in scope query parameter -- can't parse which type of access is being requested for the resource")
			}
		}

		access := ScopeAccess{
			Type: "repository",
			Name: name,
			Actions: ActionsAllowed{
				Allowed: allowed,
			},
		}

		//fmt.Println(access)
		accesses = append(accesses, access)
	}

	return accesses, nil
}

// todo right now I'm passing around slices everywhere, but not actually using them (only the first entry)
func ScopeToResponse(access []ScopeAccess) string {
	// we're just gonna assume that this ScopeAccess object is fully formed //todo don't make this assumption
	var buffer bytes.Buffer

	scopeAccess := access[0]

	if strings.EqualFold(scopeAccess.Type, "repository") {

		pushAllowed := false
		if contains(scopeAccess.Actions.Allowed, "push") {
			buffer.WriteString("repository:")
			buffer.WriteString(scopeAccess.Name)
			buffer.WriteString(":push")
			pushAllowed = true
		}

		if contains(scopeAccess.Actions.Allowed, "pull") {
			if pushAllowed {
				buffer.WriteString(",")
			}
			buffer.WriteString("repository:")
			buffer.WriteString(scopeAccess.Name)
			buffer.WriteString(":pull")
		}
	}

	return buffer.String()
}

// small helper to determine if a given byte, when a string, represents a digit
func isByteADigit(b byte) bool {
	bStr := string(b)
	digits := "1234567890"
	if strings.Contains(digits, bStr) {
		return true
	}
	return false
}

func contains(slice []string, term string) bool {
	for _, value := range slice {
		if value == term {
			return true
		}
	}
	return false
}
