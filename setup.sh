#!/bin/bash
echo -e "\e[31m\e[1m
    ______              ______               __         
   / ____/__    __     / ____/________ _____/ /__  _____
  / /  __/ /___/ /_   / / __/ ___/ __ \`/ __  / _ \/ ___/
 / /__/_  __/_  __/  / /_/ / /  / /_/ / /_/ /  __/ /    
 \____//_/   /_/     \____/_/   \__,_/\__,_/\___/_/     
                                                              
    \e[0mC++ Grader (Wizard Setup for Judge_Server)
    GoLang Version -- Made in TH/A.
    Code by... @_ngixx's (TheNongice Wasawat)
    Contacts: ngixx@ngixx.in.th

"

git clone https://github.com/TheNongice/go-grader ~/go-grader
echo -e "\n"

cd ~/go-grader
rm -rf .git
mkdir problem
mkdir runner
mkdir runner/isolate_logs
mkdir runner/temp_code
mkdir runner/temp_problem
mkdir runner/temp_code/output
mkdir runner/temp_code/cpp
mkdir runner/temp_code/cpp/output

CURRENT_DIR=$(pwd)
echo -e "\e[92m[WIZARD]\e[0m Now we're woking at: \e[92m$CURRENT_DIR\e[0m"

echo "DIR_GRADER_PATH=$CURRENT_DIR/" >> .env
echo "#################################"
echo "Welcome to wizard set-up for Simple Code Judging"
echo "Installations Guide:"
echo "1) Install GoLang with Version 1.23.x"
echo "2) Install Isolate (Sandbox Executing)"
echo " 2.1) Please check where's isolate directory is install"
echo " 2.2) The default of isolate directory path usually be:"
echo "   - /var/local/lib/isolate/"
echo "   - /usr/local/etc/isolate/"
echo "   Please manual check by yourself again :)"
echo "3) Check your .env files!"
echo "#################################"

echo -e "What's your current isolate path has installed? \e[93m"
read ISOLATE_PATH
echo "ISOLATE_PATH=$ISOLATE_PATH" >> .env

echo -e "\e[0mPlease enter username for admin: \e[93m"
read MONIT_USER
echo "MONIT_USER=$MONIT_USER" >> .env

echo -e "\e[0mPlease enter password for admin: \e[93m"
read -s MONIT_PASS
echo "MONIT_PASS=$MONIT_PASS" >> .env

echo -e "\e[92mCongratulations! This Script is install succesfully! ;)\e[0m\n"