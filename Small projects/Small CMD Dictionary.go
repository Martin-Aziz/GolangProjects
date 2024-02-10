package main

import (
	"fmt"
)

// Define a map to store word definitions
var dictionary = map[string]string{
	"algorithm":        "a step-by-step procedure for solving a problem or accomplishing a task, often used in computer science and mathematics.",
	"API":              "an application programming interface, a set of rules and protocols that allows different software applications to communicate with each other.",
	"cloud":            "a network of remote servers hosted on the Internet to store, manage, and process data, instead of on a local server or personal computer.",
	"database":         "a structured collection of data stored electronically in a computer system, typically organized to facilitate rapid search and retrieval of information.",
	"encryption":       "the process of converting data into a code to prevent unauthorized access, often used to secure sensitive information transmitted over networks.",
	"firewall":         "a security system designed to prevent unauthorized access to or from a private network by monitoring and controlling incoming and outgoing network traffic.",
	"HTML":             "HyperText Markup Language, the standard markup language for creating web pages and web applications.",
	"JavaScript":       "a high-level programming language commonly used for creating interactive effects within web browsers.",
	"server":           "a computer or software program that provides services or resources to other computers or programs on a network.",
	"virtualization":   "the process of creating a virtual version of something, such as a server, storage device, or network resource, to optimize resources and increase efficiency.",
	"VPN":              "a virtual private network, a secure connection between two or more devices or networks over the Internet, often used to protect data privacy and security.",
	"XML":              "eXtensible Markup Language, a markup language similar to HTML but designed to transport and store data, often used for representing structured data in web services and databases.",
	"agile":            "a project management methodology that emphasizes iterative development, collaboration, and flexibility to adapt to changing requirements and deliver value quickly.",
	"big data":         "a term that refers to large and complex datasets that are difficult to process using traditional data processing applications, often characterized by volume, velocity, and variety.",
	"cybersecurity":    "the practice of protecting computer systems, networks, and data from security breaches, unauthorized access, and cyber attacks.",
	"devops":           "a software development methodology that combines software development (Dev) with information technology operations (Ops), emphasizing collaboration, automation, and continuous delivery.",
	"machine learning": "a subset of artificial intelligence that uses statistical techniques to enable computers to learn from and make predictions or decisions based on data, without being explicitly programmed.",
	"networking":       "the practice of connecting computers and other devices to share resources, exchange information, and communicate with each other, often through wired or wireless networks.",
	"programming":      "the process of writing instructions for a computer to perform specific tasks, often using programming languages such as C, Java, Python, and Ruby.",
	"scalability":      "the ability of a system, network, or process to handle a growing amount of work or users by adding resources or modifying existing infrastructure without impacting performance or functionality.",
	"UI/UX":            "user interface (UI) and user experience (UX), the design and usability aspects of software applications and websites that focus on enhancing user satisfaction and interaction.",
	"version control":  "the management of changes to documents, programs, and other information over time, often using version control systems such as Git and Subversion to track revisions and facilitate collaboration.",
	"wireframe":        "a visual representation of the layout and structure of a web page or application, used to plan and design user interfaces and interactions before development begins.",
}

func main() {
	// Prompt user to enter a word
	fmt.Print("Enter a word: ")
	var word string
	fmt.Scanln(&word)

	// Lookup the word in the dictionary
	definition, found := dictionary[word]
	if found {
		fmt.Printf("Definition of %s: %s\n", word, definition)
	} else {
		fmt.Printf("Sorry, the word '%s' is not in the dictionary.\n", word)
	}
}
