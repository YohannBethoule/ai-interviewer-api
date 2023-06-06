package domain

import "fmt"

var promptJobSummaryFormat = "I need you to summarize the following job description in 100 words maximum, only emphasizing on the company context, the required skills and the ideal profile. I don't need you to write formal salutation. Here's the job description : %s"
var promptFormat = "Tu es un recruteur qui me reçoit en entretien pour le poste de %s.%s J'ai %d années d'expérience dans le domaine. Tu dois me poser une question, et attendre ma réponse pour la critiquer, avant d'en poser une nouvelle. Voici la fiche du poste pour lequel je candidate : %s"
var promptInterviewerPositionFormat = " Tu occupes le poste de %s dans cette entreprise."

func GenerateBasePrompt(jobTitle string, yearsOfExperience uint, interviewerPosition string, jobDescription string) string {
	var interviewPositionPrompt string
	if interviewerPosition != "" {
		interviewPositionPrompt = fmt.Sprintf(promptInterviewerPositionFormat, interviewerPosition)
	} else {
		interviewPositionPrompt = ""
	}
	return fmt.Sprintf(promptFormat, jobTitle, interviewPositionPrompt, yearsOfExperience, jobDescription)
}

func GenerateJobSummaryPrompt(jobDescription string) string {
	return fmt.Sprintf(promptJobSummaryFormat, jobDescription)
}
