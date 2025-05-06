# 🌐 Portfolio App

This is my personal **DevOps & Cloud learning project**, designed to serve as both a portfolio site and a platform for experimenting with modern infrastructure, automation, and deployment tools.

## 🚀 Project Overview

This project is hosted on **AWS EC2**, containerized with **Docker**, automated via **Ansible**, and deployed using **GitHub Actions** as part of a fully functional **CI/CD pipeline**.

The primary goal is to **learn by doing** — to continuously improve my skills in:

- 🛠️ Infrastructure as Code (IaC)
- 🔁 Automation and configuration management
- ☁️ Cloud deployment with AWS
- 🚀 CI/CD best practices

## 🧰 Tech Stack

- **AWS EC2** – Cloud instance hosting
- **Docker** – App containerization
- **Ansible** – Server provisioning and automation
- **GitHub Actions** – CI/CD pipelines
- **Go** – Backend web server
- **HTML/CSS** – Static front-end
- **Nginx** – Reverse proxy

## 📁 Project Structure
.
├── ansible/ # Infrastructure automation (roles, playbooks, inventory)
│ ├── playbooks/
│ └── roles/ # Includes docker, nginx, golang setup etc.
├── app/ # Golang-based web app and static files
│ ├── static/ # CSS and media (e.g. gif)
│ └── templates/ # HTML templates
├── db/ # Database init script and data volume
├── nginx/ # Nginx configuration and Dockerfile
├── docker-compose.yml
├── Dockerfile # App container definition
└── README.md

## 📦 Key Features

- Automated provisioning of an EC2 instance using Ansible
- Multi-container setup with Docker Compose
- Continuous deployment pipeline triggered by GitHub Actions
- Reverse proxy configuration with Nginx
- Simple Go-based web app with static content

## 📈 Project Status

> 🛠️ **In Progress** – I'm actively developing and expanding this project as I explore new tools, cloud services, and best practices in the DevOps space.


## 🔗 Links

-


## 🙌 Contributions

This is a personal learning project, but if you have suggestions or ideas for improvements, feel free to open an issue or fork the repo!
