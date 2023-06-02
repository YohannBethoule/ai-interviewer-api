package domain

import "fmt"

var promptFormat = "Tu es un recruteur%s qui me reçoit en entretien pour le poste de %s. J'ai %d années d'expérience dans le domaine. Tu dois me poser une question, et attendre ma réponse pour la critiquer, avant d'en poser une nouvelle. Voici la fiche du poste pour lequel je candidate : %s"
var promptInterviewerPositionFormat = "Tu occupes le poste de %s dans cette entreprise."

func GenerateBasePrompt(jobTitle string, yearsOfExperience uint, interviewerPosition string, jobDescription string) string {
	var prompt string
	if interviewerPosition != "" {
		prompt = fmt.Sprintf(promptFormat, "", jobTitle, yearsOfExperience, jobDescription)
	} else {
		prompt = fmt.Sprintf(promptFormat, fmt.Sprintf(promptInterviewerPositionFormat, interviewerPosition), jobTitle, yearsOfExperience, jobDescription)
	}
	return prompt
}
