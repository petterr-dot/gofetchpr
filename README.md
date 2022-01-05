## gofetchpr 

a simple cmd line tool to list open pull request in your repositories

#### built with 

- [go](https://go.dev/)
- [github api](https://docs.github.com/en/rest)

### prerequisites

##### Go

Download the Go installer from https://go.dev/doc/install

1. Extract the archive you downloaded into /usr/local, creating a Go tree in /usr/local/go.

   **Important:** This step will remove a previous installation at /usr/local/go, if any, prior to extracting. Please back up any data before proceeding.

   For example, run the following as root or through `sudo`:

   ```
   rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz
   ```

2. Add /usr/local/go/bin to the PATH environment variable.

   You can do this by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):

   ```
   export PATH=$PATH:/usr/local/go/bin
   ```

   **Note:** Changes made to a profile file may not apply until the next time you log into your computer. To apply the changes immediately, just run the shell commands directly or execute them from the profile using a command such as `source $HOME/.profile`.

3. Verify that you've installed Go by opening a command prompt and typing the following command:

   ```
   $ go version
   ```

4. Confirm that the command prints the installed version of Go.



##### Github API Key

1. Go to [github.com](https://github.com) and access your settings (Click your avatar in the top right corner)
2. Go to Developer settings and Personal access tokens
3. Press Generate new token. 
   1.  Add a note so you know what this Key is for, i.e "fetch pull requests"
   2.  Set expiration to something appropriate, i.e 90 days
   3. Under "Select scopes" check Repo
   4. Leave everything else default and press Generate Token at bottom of page
4. The generated token will be shown now. Make sure to save it somewhere as you wont be able to see it again. 
5. Next to the generated token press the Configure SSO button and "Authorize" to enable access to your Organization



### getting started 

1. Clone the repository 

   ```
   git clone git@github.com:petterr-dot/gofetchpr.git
   ```

2. Inside the repository there is a file called settings.go which contains `var settingsJson`. Modify this variable to contain your api_key, organization, and repositories.

2. Build the binaries with

   ```
   go install
   ```



3. Type `gofetchpr` from anywhere to execute the application

​		




 

