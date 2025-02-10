# Quackly

A companion-project that allows basic REST API deployment and basic data storage, as well as email notifications through an SMTP server, for submitting bug reports, questions, and comments. The tool returns an appropriate message for the request. If a question is submitted that is similar enough to another question, through eventually through RAG (Retrieval-Augmented Generation) we could parse through questions and author-provided answers to

## Intended Use-Case

Let's imagine that you are creating an app that would contain a user page to submit bugs, ask questions, and accept feedback in general. Sending this to an API that can handle all these requests could serve useful as a general extension to multiple apps for handling feedback and relaying these requests cleanly through email. A process could look as follows:

1. User submits a question through the app
2. The question is relayed to Quackly through an API request
3. Quackly will store the question asked (no user information is stored!)
4. [_FUTURE DEVELOPMENT_] If the RAG cannot find the question-answer combination, the question is send via email as a new user question to be answered
5. The question-answer combination will be added to the chosen data storage
6. The API request sends back one of the following:
   - The answer if it is found
   - A message indicating that the question will be processed is returned
