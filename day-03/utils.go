package daythree

const upperCaseOffset = 64
const upperCaseCorrection = 26
const lowerCaseOffset = 96 
const lowerCaseBound = 97

func pointsForChar(c rune) int {
  if int(c) < lowerCaseBound {
    return int(c) - upperCaseOffset + upperCaseCorrection
  }

  return int(c) - lowerCaseOffset 
}
