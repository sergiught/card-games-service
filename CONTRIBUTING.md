# Contributing Guidelines

Thank you for your interest in contributing to this project! To ensure a smooth collaboration process, here are outlined
some guidelines that all contributors should follow.

## Prerequisites

Before you begin contributing to the project, please ensure you have the following software installed on your system:

- [pre-commit](https://pre-commit.com/#install): This project utilizes pre-commit hooks to ensure code standards and 
minimize issues. Please install the pre-commit tool and set up the hooks by running:

```shell
pre-commit install --hook-type pre-commit --hook-type commit-msg 
```

## Commit Message Format

This project follows the [Conventional Commits](https://www.conventionalcommits.org/) specification for commit messages.
This is a lightweight convention that uses structured commit messages to make the commit history easier to understand
and manage. Here are the key rules you should follow:

- Commit messages should be structured as follows:

```text
<type>[optional scope]: <description>

[optional body]

[optional footer]
```

Where **type** is one of **feat**, **fix**, **docs**, **style**, **refactor**, **test**, **chore**, and **perf** to
represent the purpose of the commit.

- Examples of Conventional Commits:
  - feat(blog): add comment section under blog post
  - fix(api): handle null error when user data is missing
  - docs(readme): update installation instructions

Please ensure that your commit messages follow this format. It helps in automated changelog generation and streamlines
the release process.
