# THIS PROJECT IS IN ITS EARLY STAGES

## This readme is just a design documentation for now

# godo

Simple job remoting tool meant as an alternative to buildbot.

## Name

godo is a shit name, you can help by finding a better one. For now I can't be bothered to change it.

## Core principles

### Master, Slave and Servant

#### Master

The master has numerous jobs defined by simple configuration files. It also holds a list of slaves and servants that he will contact to provide them with said jobs.

#### Slave

The slave will be contacted by the master and provided with a job to run. If the tags of the slave allow him to run said job he will do so and send back the result of the execution to the master. If the tags forbid him from running the job, he will notify the master who will attempt to pick a better suited slave.

#### Servant

The servant is a specific kind of slave that also acts as a master. Like a master it has a list of slaves it will try ton contact, but like the slave he also has tags and will try to run any job that's provided to him. If he can't run the job he's been provided with, it will try to pass it on its own slaves. With that specific configuration you can have the slave be the master of another slave and sending it work to do.

### Jobs are not scripts

A job is defined by a name, an optionnal list of tags and holds a one line expression that will be evaluated by the shell of the slave or servant assigned to it.

Unlike pieces of software like _buildbot_, this means that you don't have to delve into a specific API to configure the tasks you want to execute on your remote machines, you can express them in any shell scripting language your slave can handle or rely on custom script or executables.

### Tags drive help decide what to do

Job tags, or tags for short, are short strings used to determine if a slave has the ability to execute a proposed job. These tags usually are named after a piece of software, the OS type, or configuration present on the slave machine and/or needed to run a job. You can think of them as a way to detect the dependencies for a job.

These tags can be written by hand or generated automatically from a list of known frequently used pieces of software.

If no tags are specified on the slave, it will try to execute every job that's provided to him. If no tags are specified on the job, every slave it is provided to will try to run it.

This behavior is intended to simplify configuration when you know for sure that the needed dependencies are present on your slave machines and you don't want to bother with the tagging system.

## Installation

    TODO

## Usage

### Master, servant and slave configuration and control

#### Create config files

    godo create master [options]

    godo create servant [options]

    godo create slave [options]

#### Add a slave to the current master or servant configuration

    godo config add slave [host:port]

#### Open the local configuration file with the default text editor

    godo config edit

#### General controls

    godo start [options]

    godo restart [options]

    godo stop

    godo status

### Job creation, configuration and control

    godo job create <name> [options]

    godo job delete <name> [option]

    godo job start <name>

    godo job edit <name>

## Configuration

## File locations

By default, godo will look for its configuration files, starting with the main godo.conf file, at the following locations:

    /etc/godo/godo.conf

    /etc/godo/jobs.d

,

    /usr/local/share/godo/godo.conf

    /usr/local/share/godo/jobs.d

and

    ~/.config/godo/godo.conf
    
    ~/.config/godo/jobs.d

## TODO

- Use flag package to parse [options], the rest I should probablu handle manually
- Use TOML
