<p align="center">
    <a href="#" target="_blank" rel="noopener">
        <img src="/.github/media/banner.png" alt="litecart - shopping-cart in 1 file" />
    </a>
</p>


<a href="https://github.com/shurco/litecart/releases"><img src="https://img.shields.io/github/v/release/shurco/litecart?sort=semver&label=Release&color=651FFF"></a>
<a href="https://goreportcard.com/report/github.com/shurco/litecart"><img src="https://goreportcard.com/badge/github.com/shurco/litecart"></a>
<a href="https://www.codefactor.io/repository/github/shurco/litecart"><img src="https://www.codefactor.io/repository/github/shurco/litecart/badge" alt="CodeFactor" /></a>
<a href="https://github.com/shurco/litecart/actions/workflows/release.yml"><img src="https://github.com/shurco/litecart/actions/workflows/release.yml/badge.svg"></a>
<a href="https://github.com/shurco/litecart/blob/master/LICENSE"><img src="https://img.shields.io/badge/License-MIT-yellow.svg"></a>


## 🛒&nbsp;&nbsp;What is litecart?

Litecart is an open source shopping-cart in 1 file of embedded database (SQLite), convenient dashboard UI and simple site.

> ⚠️&nbsp;&nbsp;Current major version is zero (`v0.x.x`) to accommodate rapid development and fast iteration while getting early feedback from users. Please keep in mind that litecart is still under active development and therefore full backward compatibility is not guaranteed before reaching v1.0.0.

![Example](/.github/media/demo.gif)

## 🏆&nbsp;&nbsp;Features

- [x] 🚀 **Simple and Fast**: Enjoy a one-click installation process that gets your store up and running quickly, saving you time and effort. 
- [x] 💰 **Support for Popular Payment Systems**: Accept payments seamlessly with support for popular payment systems, ensuring a smooth checkout experience for your customers. 
- [x] 🔑 **Sell Files and License Keys**: Whether you're selling digital files or license keys, litecart has you covered, providing flexibility in the types of products you can offer.
- [x] ⚙️ **Lightweight and Efficient**: litecart utilizes SQLite as its embedded database, eliminating the need for heavy databases like MySQL, PostgreSQL, or MongoDB. This results in a lightweight website that performs exceptionally well.
- [x] ☁️ **Easily Customizable**: Modify and customize your litecart website effortlessly to match your branding and unique requirements, making it truly your own.
- [x] 🧞‍♂️ **Convenient Administration Panel**: With a user-friendly dashboard UI, litecart offers a hassle-free administration panel, allowing you to manage your store, inventory, and orders with ease. 
- [x] ⚡️ **Hardware Compatibility**: Whether you're running litecart on a powerful server or a modest hardware setup, rest assured that it will work seamlessly, providing a consistent shopping experience for your customers.
- [x] 🔒 **Built-in HTTPS Support**: Prioritizing security, litecart comes with built-in support for HTTPS, ensuring the safety of your customers' data.




## ⬇️&nbsp;&nbsp;Installation

`litecart` is engineered for easy installation and operation, requiring just a single command from your terminal. Besides the conventional installation method, `litecart` can also be set up and operated via HomeBrew, Docker, or any other container orchestration tools like Docker Compose, Docker Swarm, Rancher, or Kubernetes.

#### <img width="20" src="/.github/media/platforms/apple.svg">&nbsp;Install on macOS
The fastest method to install `litecart` on macOS involves using Homebrew. This will install the command-line tools and the `litecart` server as a combined executable. If you don't utilize Homebrew, adhere to the Linux instructions below for `litecart` installation.
```shell
brew install shurco/tap/litecart
```

Alternately, you can configure the tap and install the package separately:
``` shell
$ brew tap shurco/tap
$ brew install litecart
```


#### <img width="20" src="/.github/media/platforms/linux.svg">&nbsp;Install on Linux 
The most straightforward and recommended method to start using `litecart` on Unix operating systems involves installing and utilizing the `litecart` command-line tool. Execute the given command in your terminal and adhere to the instructions displayed on the screen.

```bash
curl -L https://raw.githubusercontent.com/shurco/litecart/main/scripts/install | sh
```

#### <img width="20" src="/.github/media/platforms/windows.svg">&nbsp;Install on Windows
The simplest and most recommended method to start using `litecart` on Windows is by installing and utilizing the `litecart` command-line tool. Execute the given command in your terminal and adhere to the instructions displayed on the screen.
```bash
curl -L https://raw.githubusercontent.com/shurco/litecart/main/scripts/install | sh
```
or download and unzip the [latest version](https://github.com/shurco/litecart/releases/latest) for Windows.


#### <img width="20" src="/.github/media/platforms/docker.svg">&nbsp;Run using Docker
Docker enables the management and operation of a `litecart` instance without requiring the installation of any command-line tools. The `litecart` Docker container includes all necessary command-line tools  or even for server execution.

For [Docker Hub](https://hub.docker.com/r/shurco/litecart):
```bash
docker run \
  -v ./lc_base:/lc_base \
  -v ./lc_digitals:/lc_digitals \
  -v ./lc_uploads:/lc_uploads \
  -v ./site:/site \
  --rm shurco/litecart:latest init

docker run \
  --name litecart \
  --restart unless-stopped \
  -p '8080:8080' \
  -v ./lc_base:/lc_base \
  -v ./lc_digitals:/lc_digitals \
  -v ./lc_uploads:/lc_uploads \
  -v ./site:/site \
  shurco/litecart:latest
```
or if use [Github Packages Hub](https://github.com/shurco/litecart/pkgs/container/litecart):

```bash
docker run \
  -v ./lc_base:/lc_base \
  -v ./lc_digitals:/lc_digitals \
  -v ./lc_uploads:/lc_uploads \
  -v ./site:/site \
  --rm ghcr.io/shurco/litecart:latest init

docker run \
  --name litecart \
  --restart unless-stopped \
  -p '8080:8080' \
  -v ./lc_base:/lc_base \
  -v ./lc_digitals:/lc_digitals \
  -v ./lc_uploads:/lc_uploads \
  -v ./site:/site \
  ghcr.io/shurco/litecart:latest
```

## ⬇️&nbsp;&nbsp;Updating
> ⚠️&nbsp;&nbsp;Before any update, be sure to make a backup of the *./lc_base* folder and the *./site* folder.

#### Update on macOS / Linux / Windows
The easiest way to update `litecart` to the latest version is to execute the command:

```bash
./litecart update
```

If there were changes in the database structure during the update, it is necessary to perform migration. To do this, you need to run the command from the `litecart` folder:
```bash
./litecart migrate
```


#### <img width="20" src="/.github/media/platforms/docker.svg">&nbsp; Update using Docker
Our mantra is to make updating a seamless experience. Simply download the new image and launch the container as you normally would. For example, if use [Docker Hub](https://hub.docker.com/r/shurco/litecart):

```bash
docker stop litecart
docker pull shurco/litecart:latest # download new image
docker rename litecart litecart-backup # do image backup
docker run \
  --name litecart \
  --restart unless-stopped \
  -p '8080:8080' \
  -v ./lc_base:/lc_base \
  -v ./lc_digitals:/lc_digitals \
  -v ./lc_uploads:/lc_uploads \
  -v ./site:/site \
  shurco/litecart:latest
```

If there were changes in the database structure during the update, it is necessary to perform migration. To do this, you need to run the command from the `litecart` folder:
```bash
docker run \
-v ./lc_base:/lc_base \
-v ./site:/site \
--rm shurco/litecart migrate
```

## 🚀&nbsp;&nbsp;Getting started
Getting started with `litecart` is as easy as starting up the `litecart` server

Default run for Linux/macOS:
```bash
./litecart serve
```

For Windows:
```
litecart.exe serve
```

When launched for the first time, necessary folders will be created in the directory with the executable file. The default links for access are:  
- [http://localhost:8080](http://localhost:8080) - website  
- [http://localhost:8080/_/](http://localhost:8080/_/) - control panel  

If you need to run on a different port, use the flag `--http`:
```
./litecart serve --http 0.0.0.0:8088
```

> ⚠️&nbsp;&nbsp; Ports <= 1024 are privileged ports. You can't use them unless you're root or have the explicit permission to use them. See this answer for an explanation or wikipedia or something you trust more. Use:
**sudo setcap 'cap_net_bind_service=+ep' /path_to/litecart**
> 

## 📚&nbsp;&nbsp;Commands
Usage:
```
./litecart [command] [flags]
```

Available commands:
```
init        Creating the basic structure
migrate     Migrate on the latest version of database schema
serve       Starts the web server (default to 0.0.0.0:8080)
update      Updating the application to the latest version
```

Global flags `./litecart [flags]`:
```
-h, --help      help for litecart
-v, --version   version for litecart
```

Serve flags `./litecart serve [flags]`:
```
--http string    server address (default "0.0.0.0:8080")
--https string   https server address (auto TLS)
--no-site        disable create site
```


## 🗺️&nbsp;&nbsp;ToDo
`litecart` has a [roadmap](https://github.com/users/shurco/projects/2) and I try to work on issues in specific order and such PRs often come in out of nowhere and skew all initial planning with tedious back-and-forth communication.

- [x] Product in the form of files
- [x] Product in the form of license keys
- [ ] Product returned via API to another site (example license keys)
- [x] Payment Stripe
- [ ] Payment PayPal
- [ ] Payment Square
- [ ] Payment Adyen
- [ ] Payment Checkout
- [ ] Support for payment using crypto
- [ ] Support WebHook


## 👍&nbsp;&nbsp;Contribute

If you want to say **thank you** and/or support the active development of `litecart`:

1. Add a [GitHub Star](https://github.com/shurco/litecart/stargazers) to the project.
2. Tweet about the project [on your Twitter](https://twitter.com/intent/tweet?text=%F0%9F%9B%92%20litecart%20-%20shopping-cart%20in%201%20file%20on%20%23Go%20https%3A%2F%2Fgithub.com%2Fshurco%2Flitecart).
3. Write a review or tutorial on [Medium](https://medium.com/), [Dev.to](https://dev.to/) or personal blog.
4. Support the project by donating a [cup of coffee](https://github.com/sponsors/shurco).

You can learn more about how you can contribute to this project in the [contribution guide](https://github.com/shurco/litecart/blob/master/.github/CONTRIBUTING.md).
