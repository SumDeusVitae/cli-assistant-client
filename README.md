<div align="center">
  <img src="https://github.com/SumDeusVitae/cli-assistant-client/blob/main/assistantHeader.png" />
</div>

# CLI Assistant
## Table of Contents
- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Example](#example)
- [Uninstall](#uninstall)
- [License](#license)

## Introduction
CLI Assistant is a command-line tool that lets you ask questions and get answers directly in the terminal. It's perfect for situations where you don’t have access to a web browser, but still need quick answers or help solving problems. Whether you need a quick reminder, a bit of troubleshooting, or just some general knowledge, CLI Assistant has you covered—all from the comfort of your command line.

## Installation
### 1. Install Go 1.22 or later

The Boot.dev CLI requires a Golang installation, and only works on Linux and Mac. If you're on Windows, you'll need to use WSL. Make sure you install go in your Linux/WSL terminal, not your Windows terminal/UI. There are two options:

**Option 1**: [The webi installer](https://webinstall.dev/golang/) is the simplest way for most people. Just run this in your terminal:

```bash
curl -sS https://webi.sh/golang | sh
```

_Read the output of the command and follow any instructions._

**Option 2**: Use the [official installation instructions](https://go.dev/doc/install).

Run `go version` on your command line to make sure the installation worked. If it did, _move on to step 2_.

**Optional troubleshooting:**

- If you already had Go installed with webi, you should be able to run the same webi command to update it.
- If you already had a version of Go installed a different way, you can use `which go` to find out where it is installed, and remove the old version manually.
- If you're getting a "command not found" error after installation, it's most likely because the directory containing the `go` program isn't in your [`PATH`](https://opensource.com/article/17/6/set-path-linux). You need to add the directory to your `PATH` by modifying your shell's configuration file. First, you need to know _where_ the `go` command was installed. It might be in:

- `~/.local/opt/go/bin` (webi)
- `/usr/local/go/bin` (official installation)
- Somewhere else?

You can ensure it exists by attempting to run `go` using its full filepath. For example, if you think it's in `~/.local/opt/go/bin`, you can run `~/.local/opt/go/bin/go version`. If that works, then you just need to add `~/.local/opt/go/bin` to your `PATH` and reload your shell:

```bash
# For Linux/WSL
echo 'export PATH=$PATH:$HOME/.local/opt/go/bin' >> ~/.bashrc
# next, reload your shell configuration
source ~/.bashrc
```

```bash
# For Mac OS
echo 'export PATH=$PATH:$HOME/.local/opt/go/bin' >> ~/.zshrc
# next, reload your shell configuration
source ~/.zshrc
```

### 2. Installing CLI ASSISTANT
This command will download, build, and install the `cli-assistant-client` command into your Go toolchain's `bin` directory. Go ahead and run it:
```bash
go install github.com/SumDeusVitae/cli-assistant-client@latest
```
After this, you can rename the binary:
```bash
mv $(go env GOPATH)/bin/cli-assistant-client $(go env GOPATH)/bin/qs
```
Now, you can run your app using qs instead of cli-assistant-client.
Run this command to check:
```bash
qs version
```
If it didn't work try:
```bash
cli-assistant-client version
```
If it works instead that means that app installed correctly but you didn't rename binary correctly in your environment PATH


## Usage
To run tool, use following commands
For new users please register:
```bash
qs register
```
If you have an account you can simply login:
```bash
qs login
```

after you registered or logged in use:
```bash
qs q <your question here>
```


## Commands 
- `help`:  Displays a help message
- `register`: Registers new user
- `login`:  Login as existent user
- `q <question>`:  Asks AI question 
- `whoami`:  Checks if you logged in
- `helth`:  Checking server status
- `version`:  Checks current version
- `env`:  Shows saved environmental variables
- `update`:  Updates CLI to the latest version

## Example
<pre>
user@DESKTOP:~/workspace$ qs q How do I add tag to git repo?
Working on respond...

Reply: 
Adding a tag to a Git repository is a straightforward process. Tags are used to mark specific points in your repository’s history, often used to denote releases or important milestones.

Here’s how to add a tag:

### 1. Create a Tag

You can create two types of tags in Git: lightweight and annotated.

#### Lightweight Tag

A lightweight tag is simply a pointer to a specific commit. It does not contain additional information such as the tagger's name or date.

```bash
git tag <tagname>
```

#### Annotated Tag

Annotated tags are recommended for most cases, as they store additional metadata such as the tagger's name, email, date, and a message.

```bash
git tag -a <tagname> -m "Your message here"
```

### 2. List Tags

To see the tags you've created, you can list them with:

```bash
git tag
```

### 3. Push Tags to Remote

By default, tags are not pushed to remote repositories when you run `git push`. To push a specific tag to the remote repository, use:

```bash
git push origin <tagname>
```

If you want to push all tags at once, you can use:

```bash
git push --tags
```

### 4. Checkout a Tag

To check out a specific tag (detached HEAD state), you can use the following command:

```bash
git checkout <tagname>
```

### Example

Here’s a quick example of the entire process:

1. Create a new tag:

   ```bash
   git tag -a v1.0 -m "Release version 1.0"
   ```

2. List tags to confirm:

   ```bash
   git tag
   ```

3. Push the tag to the remote repository:

   ```bash
   git push origin v1.0
   ```


### Conclusion

That’s it! You have successfully created and pushed a tag to your Git repository. Tags are a great way to mark significant points in your project’s history, making it easier to manage releases and versions.
</pre>

## Uninstall
To uninstall CLI Assistant, follow the steps :

  Locate the binary: Find the binary in your GOBIN or GOPATH/bin directory.
  Remove the binary: Delete it from the terminal:

```bash
rm $(go env GOPATH)/bin/qs
```
  Remove the source code (optional):
```bash
rm -rf $(go env GOPATH)/src/github.com/SumDeusVitae/cli-assistant-client
```


## License
This project is licensed under the [MIT License](LICENSE).
