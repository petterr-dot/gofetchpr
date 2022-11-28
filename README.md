## gofetchpr 

a simple cmd line tool to list open pull request in your repositories

#### built with 

- [go](https://go.dev/)
- [github api](https://docs.github.com/en/rest)

### prerequisites

##### Go

Download and install Go (if you dont have it already) https://go.dev/doc/install



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
   Every time you update settings.go (adding/removing repositories) you need to build the binaries again with this command



3. Type `gofetchpr` from anywhere to execute the application

​		




 

