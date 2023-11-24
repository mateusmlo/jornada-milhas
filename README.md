[![Go](https://github.com/mateusmlo/jornada-milhas/actions/workflows/go.yml/badge.svg)](https://github.com/mateusmlo/jornada-milhas/actions/workflows/go.yml)
[![Docker Image CI](https://github.com/mateusmlo/jornada-milhas/actions/workflows/docker-image.yml/badge.svg)](https://github.com/mateusmlo/jornada-milhas/actions/workflows/docker-image.yml)
<div id="top"></div>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->


### ⚠️ WORK IN PROGRESS
Some things may not be working as intended or just not working at all, I'm doing my best 🥲

<!-- ABOUT THE PROJECT -->
## About The Project

My entry to the 7th edition of Alura's back-end challenge, which consists of an API from a travel agency where users can search for possible travel locations, write reviews and attach pictures. For this project I decided to use my second language, which is Go! It has been very fun and challenging to code, since it was also my first time using the Gin framework and also Uber's absolutely magical dependency injection system Fx, which made me think about Go packages a bit similar to what I'm used to in NestJS. I took heavy inspiration from several other projects and a few videos from awesome people, check out the <a href="#acknowledgments">Acknowledgments</a> section for them! But as usual I've added my own secret spice to the code as well 😉


<p align="right">(<a href="#top">back to top</a>)</p>

### Built With

* [Go v1.21](https://go.dev/)
* [Gin](https://gin-gonic.com/)
* [Fx](https://uber-go.github.io/fx/)
* [PostgreSQL](https://www.postgresql.org/)
* [Docker](https://www.docker.com/)
* [Air](https://github.com/cosmtrek/air)
* ...and more to come!

<p align="right">(<a href="#top">back to top</a>)</p>

### TODO
- [ ] Add refresh tokens
- [ ] Deploy on cloud

<!-- GETTING STARTED -->
## Getting Started

Getting this API up and running is extremely fast and simple! There are a few ways depending on what you want to do: running locally and debug, or just test the endpoints.

### Prerequisites

This project was built on Go v1.21.0 and makes heavy use of Docker so both are must have if you plan on running locally. If not, it is completely possible to just spin a docker-compose container and have fun.
First of all, provide your values of choice to a `.env` file which should be a copy of the `.env.example`.

  ```sh
  go mod download
  ```

### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Run the project's external dependencies
   ```sh
   docker-compose up -d
   ```
3. To run the API locally:
   ```sh
   air ./cmd/jornada-milhas
   ```

You might need to comment the project's image under the `docker-compose.yml` file if you want to debug.

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Mateus Oliveira - [LinkedIn](https://www.linkedin.com/in/mateusmlo/) - mateus.mlo95@gmail.com

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
## Acknowledgments
Awesome projects and content that helped me figure things out for this little beast of a project!

Useful videos:
* [Trying Another Way... (Dependency Injection)](https://www.youtube.com/watch?v=8Oosc55SKrM&t=508s)
* [Go Programming – Golang Course with Bonus Projects](https://www.youtube.com/watch?v=un6ZyFkqFKo&t=10204s)


A template for clean code using Gin and Fx. Didn't use it for this project as I thought it would be too "overkill" for the scope:
* [Clean Gin](https://github.com/dipeshdulal/clean-gin)


Below are some entries for a back-end contest held back in August here in Brazil that I've studied a lot as well:
* [Leo Vargas' entry](https://github.com/leorcvargas/rinha-go)
* [Isadora Souza's entry](https://github.com/isadoramsouza/rinha-de-backend-go/tree/main)

<p align="right">(<a href="#top">back to top</a>)</p>
