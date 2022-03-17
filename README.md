
<div align="center">

  
## ---|Welcome to my Go MCU api Project 👋|---


<img src="https://raw.githubusercontent.com/NorbertRuff/go-mcu-graphql_api/master/blob/logos/gopher2.png" width="350">


</div>

<div align="center">

This year I started learning GO and GraphQL, and wanted to give a try how they can work together,
and this little project came out of it.

You can register a user in a db using GraphQL, query, mutate the db with jwt authentication.

It uses PostgreSQL db in a docker.

This was created purely just for fun, and experimenting with the tech stacks.

Hope you like it and helps you in some way.


[![Version](https://img.shields.io/badge/version-v0.8-blue.svg)](https://img.shields.io/badge/version-v0.8-blue.svg?cacheSeconds=2592000)



•[PROJECT PHILOSOPHY](#project-philosophy) •
[TECH STACK](#tech-stack) •
[SCREENSHOTS](#screenshots) •
[CONTRIBUTING](#contributing) •
[ABOUT ME](#-about-me)
•

</div>

## Project-philosophy

Experiment with GO, learn some new things, with docker / GO / GraphQL.


## Demo

TODO

## Screenshots


<img src="https://raw.githubusercontent.com/NorbertRuff/go-mcu-graphql_api/master/blob/screenshot/screenshot.png">



<img src="https://raw.githubusercontent.com/NorbertRuff/go-mcu-graphql_api/master/blob/screenshot/screenshot2.png">



## Features
<div align="center">

<img src="https://raw.githubusercontent.com/NorbertRuff/go-mcu-graphql_api/master/blob/logos/gopher.png" width="350">

</div>

## Tech Stack

- GO
- Gorm
- gqlgen
- Postgres
- graphql
- docker

**API-s used for MCU data**

- [MCU heroku app](https://mcuapi.herokuapp.com/docs/)


**Packages used**

- github.com/99designs/gqlgen for GraphQl generator
- jwt-go
- Go chi
- gqlparser
- gorm


<div align="center">

<img src="https://raw.githubusercontent.com/NorbertRuff/go-mcu-graphql_api/master/blob/logos/gopher3.png" width="350">

</div>

## Prerequisites

- Install go
- Install docker / docker-compose
- Set up .env file


## Installation
### Install go(lang)

with [homebrew](http://mxcl.github.io/homebrew/):

```Shell
sudo brew install go
```
  
with [apt](http://packages.qa.debian.org/a/apt.html)-get):

```Shell
sudo apt-get install golang
```

[install Golang manually](https://golang.org/doc/install)
or
[compile it yourself](https://golang.org/doc/install/source)

### Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`PORT=` -> your localhost port for the qraphql playground defaults to 8080  

`DB_HOST=`  -> localhost  
`DB_PORT=`  -> 5432  
`DB_USER=`  -> anything   
`DB_PASSWORD=`  -> anything  
`DB_NAME=`-> anything  


## Run Locally

Clone the project

```bash
  git clone https://github.com/NorbertRuff/go-mcu-graphql_api
```

Go to the project directory

```bash
  cd go-mcu-graphql_api
```

Install dependencies  

```bash
go get
```

```bash
go mod tidy
```


First start Postgres server with docker compose:

```bash
 docker-compose up -d
```

Finally run the server:

```bash
go run server.go
```

Now navigate to https://localhost:8080 you can see graphiql playground and query the graphql server.


## Commands
For example:

You can create a user, this command gives back a JWT Token

```graphql
mutation createuser {
  createUser(
    variables: {
      username: "bob"
      password: "bob"
      email: "bob"
      firstname: "bob"
      lastName: "bob"
    }
  )
}
```
Login an existing user, this command gives back a JWT Token


```graphql
mutation {
    login(variables: { username: "bob", password: "bob" })
}
```

Query all the movies requires: JWT token

```graphql
query {
    movies
    {
        imdb_id
        movie_id
        title
        overview
        duration
        user {
            user_id
            username
        }
    }

}
```


You can generate resolvers and types from gql schema configured in gqlgen.yml

```bash
go run github.com/99designs/gqlgen generate 
```

or 

```bash
go generate 
```


### Compile

One great aspect of Golang is, that you can start go applications via ```go run server.go```, but also compile it to an executable with ```go build server.go```. After that you can start the compiled version which starts much faster.

Build your app and synthesize your stacks.

> Generates a `.build/` directory with the compiled files.


<div align="center">

<img src="https://raw.githubusercontent.com/NorbertRuff/go-mcu-graphql_api/master/blob/logos/gopher4.png" width="350">

</div>

## Lessons Learned

- Basic clean code
- Go project structure
- Go graphgl
- Go pointers
- Go structs, types
- Graphql schema
- Docker compose
- Sql
- Postgres
- Gorm
- Docker Compose


Give a ⭐️ if this project helped you!


# <div style="color:#f59800" align="center"> 🚀About-Me</div>

## <h2 align="center">Hi 👋, I'm Norbert</h2>
### A passionate developer from Hungary

</div>
 
- 🔭 I’m currently working on Frontend skills 🚀
- 🎓 I recently completed Codecool Hungary Full-Stack developer bootcamp.
- 👯 I’m looking to collaborate with other Developers :wink:
- 🥅 2022 Goals: Learn scss, less css, typescript, front-end frameworks.
- 💬 Ask me about anything, I am happy to help :wink:
- 🧗 I try to go beyond and push the bounds.
- ⚡ Fun fact: I love old technology :raised_hands:
  
<div align="center">

[![LinkedIn](https://img.shields.io/badge/-Norbert%20Ruff-blue?style=for-the-badge&logo=linkedin&logoColor=white&link=https://www.linkedin.com/in/ruff-norbert-6b167b204//)](https://www.linkedin.com/in/ruff-norbert-6b167b204/)
[![Gmail Badge](https://img.shields.io/badge/-ruffnorbert88@gmail.com-c14438?style=for-the-badge&logo=Gmail&logoColor=white&link=mailto:ruffnorbert88@gmail.com)](mailto:ruffnorbert88@gmail.com)
![Discord Badge](https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white)

</div>

<h2><img src="https://media.giphy.com/media/cj87CxfRtrUifF3Ryk/giphy.gif" height="25"> My Github Stats</h2>

<div align="center">

[![](https://raw.githubusercontent.com/NorbertRuff/NorbertRuff/master/profile-summary-card-output/dracula/0-profile-details.svg)](https://github.com/vn7n24fzkq/github-profile-summary-cards)

[![](https://raw.githubusercontent.com/NorbertRuff/NorbertRuff/master/profile-summary-card-output/dracula/2-most-commit-language.svg)](https://github.com/vn7n24fzkq/github-profile-summary-cards)

[![](https://raw.githubusercontent.com/NorbertRuff/NorbertRuff/master/profile-summary-card-output/dracula/3-stats.svg)](https://github.com/vn7n24fzkq/github-profile-summary-cards) [![](https://raw.githubusercontent.com/NorbertRuff/NorbertRuff/master/profile-summary-card-output/dracula/4-productive-time.svg)](https://github.com/vn7n24fzkq/github-profile-summary-cards)

</div>



## My Skill Set 👩‍💻
<!-- https://dev.to/envoy_/150-badges-for-github-pnk -->
<div align="center">  
<img src="https://www.codewars.com/users/NorbertRuff/badges/large">
</div>

<table><tr><td valign="top" width="25%">
<h2 align="center"> 💻 </h2><br>

<div align="center">  
<img src="https://img.shields.io/badge/Python-3776AB?style=flat-square&logo=python&logoColor=white" height="25">
<img src="https://img.shields.io/badge/Java-ED8B00?style=flat-square&logo=java&logoColor=white" height="25">
  
<img src="https://img.shields.io/badge/-GraphQL-E10098?style=flat-square&logo=graphql&logoColor=white" height="25">
  
<img src="https://img.shields.io/badge/shell_script-%23121011?style=flat-square&logo=shell&logoColor=white" height="25">
  
<img src="https://img.shields.io/badge/spring-%236DB33F?style=flat-square&logo=spring&logoColor=white" height="25">
   
<img src="https://img.shields.io/badge/Svelte-4A4A55?style=flat-square&logo=svelte&logoColor=FF3E00" height="25">
   
<img src="https://img.shields.io/badge/PostgreSQL-316192?style=flat-square&logo=postgresql&logoColor=white" height="25">
  
  
<img src="https://img.shields.io/badge/-jest-%23C21325?style=flat-square&logo=jest&logoColor=white" height="25">
  
  
</div>


</td><td valign="top" width="25%">

<h2 align="center"> 🌐 </h2><br>

<div align="center">  


<img src="https://img.shields.io/badge/-CSS3-1572B6?style=flat-square&logo=css3" height="25">
<img src="https://img.shields.io/badge/HTML5-E34F26?style=flat-square&logo=html5&logoColor=white" height="25">
<img src="https://img.shields.io/badge/React-20232A?style=flat-square&logo=react&logoColor=61DAFB" height="25">
<img src="https://img.shields.io/badge/Bootstrap-563D7C?style=flat-square&logo=bootstrap&logoColor=white" height="25">
<img src="https://img.shields.io/badge/JavaScript-F7DF1E?style=flat-square&logo=javascript&logoColor=black" height="25">
<img src="https://img.shields.io/badge/typescript-%23007ACC?style=flat-square&logo=typescript&logoColor=white" height="25">
<img src="https://img.shields.io/badge/Flask-000000?style=flat-square&logo=flask&logoColor=white" height="25">
<img src="https://img.shields.io/badge/nestjs-%23E0234E?style=flat-square&logo=nestjs&logoColor=white" height="25"> 
<img src="https://img.shields.io/badge/node.js-6DA55F?style=flat-square&logo=node.js&logoColor=white" height="25"> 
<img src="https://img.shields.io/badge/Thymeleaf-%23005C0F?style=flat-square&logo=Thymeleaf&logoColor=white" height="25">

</div>

</td><td valign="top" width="25%">

<h2 align="center"> ⚙ </h2><br>

<div align="center">

<img src="https://img.shields.io/badge/-Linux-black?style=flat-square&logo=Linux" height="25"> 
<img src="https://img.shields.io/badge/Windows-0078D6?style=flat-square&logo=windows&logoColor=white" height="25"> 
<img src="https://img.shields.io/badge/NPM-%23000000?style=flat-square&logo=npm&logoColor=white" height="25"> 
<img src="https://img.shields.io/badge/Ubuntu-E95420?style=flat-square&logo=ubuntu&logoColor=white" height="25">
<img src="https://img.shields.io/badge/-Git-black?style=flat-square&logo=git" height="25"> 
<img src="https://img.shields.io/badge/-GitHub-181717?style=flat-square&logo=github" height="25"> 
<img src="https://img.shields.io/badge/githubactions-%232671E5?style=flat-square&logo=githubactions&logoColor=white" height="25"> 
<img src="https://img.shields.io/badge/Markdown-000000?style=flat-square&logo=markdown&logoColor=white" height="25">
<img src="https://img.shields.io/badge/pycharm-143?style=flat-square&logo=pycharm&logoColor=white" height="25">  
<img src="https://img.shields.io/badge/Visual%20Studio%20Code-0078d7?style=flat-square&logo=visual-studio-code&logoColor=white" height="25">  
<img src="https://img.shields.io/badge/webstorm-143?style=flat-square&logo=webstorm&logoColor=white" height="25">  

</div>

</td>
</td><td valign="top" width="25%">

<h2 align="center"> 🎨 </h2><br>

<div align="center">
<img src="https://aleen42.github.io/badges/src/photoshop.svg" height="25">
<img src="https://aleen42.github.io/badges/src/illustrator.svg" height="25">
<img src="https://aleen42.github.io/badges/src/dreamweaver.svg" height="25">
<img src="https://aleen42.github.io/badges/src/flash.svg" height="25">
  
 </div>

</td>
</tr></table>  

<div align="center">

<p align="center"> <img src="https://komarev.com/ghpvc/?username=NorbertRuff&label=Profile%20views&color=0e75b6&style=flat-square" alt="prathmesh" /> </p>


</div>


## Roadmap

TODO

## Feedback

If you have any feedback, please visit the issues page


## Acknowledgements

TODO
- Logo - [Gopherize.me](https://gopherize.me/)


## Contributing

Contributions are always welcome!

See `contributing.md` for ways to get started.

Please adhere to this project's `code of conduct`.





