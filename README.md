# CryptoTracker

A Project to track cryptocurrency and gather data! It has many configurations to make gathering and pushing data togit repos easier. The steps are as follows:

Clone the repository
```
git clone https://github.com/dhillondeep/cryptotracker.git
```

Change directory to `assets/config` folder in the repo. It contains a config.yml file which let's the user modify various properties
```
cd cryptotracker/assets/config
```

Modify `config.yml` file according to your needs. After modifying, we need to come back to the root directory and build the project
```
cd ../../
make build
```

After building the project, we need to run the project. There are two options: **gather** and **monitor**. `gather` command run the program once where it gatheres data and pushes it to the repo(s). `monitor` on the other hand, runs the program continously gathering and pushing data every **interval** defined in the `config.yml` file.
```
make run ARGS="gather"
# or
make run ARGS="monitor"
# to override the data in the repo
make run ARGS="gather --override"
# specify number of commits
make run ARGS="gather --commits 15"`
```

