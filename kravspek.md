REST API


What it does:
- provides client with information about books available in a given language (based on gutenberg library)
- provides info about amount of potential readers for a certain language

Endpoints (aka the route of the URL):
- Books avalable in language
- Number of potential readers

Paths:
- /librarystats/v1/bookcount/  -> Provide user-readable guidance on how to invoke the service
- /librarystats/v1/readership/ -> Provide user-readable guidance on how to invoke the service
- /librarystats/v1/status/

Other requirements:
- Setup and use of application must be documented in a README.md file


Functions:
- One fuction to send requests, that other functions use to handle a specific request


Submission:
- Both Render URL and code repository from GitLab workspace