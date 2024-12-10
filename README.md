# Simple Code Judging
A code judging system is made from GoLang. This system is made for using with C/C++ programing languages.
Currently, This system is in prototype for using to learn how to create automation code judging system only.

*This programs is make for Debian & Ubuntu.*

# Feature
- Auto compile C/C++ source code.
- Auto judging the result from execute files.
- Check walltime, runtime to decide time out. (Using [Isolate](https://github.com/ioi/isolate))
- It can specific the wrong & correct answer.

# Setup
This program currently in alpha. I'm not make auto-install script untill it's in beta.
(Manual Setup)
1) Install GoLang (Version 1.23.x)
2) Install Isolate ([MANUAL INSTALL](https://www.ucw.cz/moe/isolate.1.html#_installation))
3) Setup `go-grader` in .env with these example!
```env
ISOLATE_PATH=/var/local/lib/isolate/
DIR_GRADER_PATH=/home/YOUR_USER/go_grader/
# Please don't forget / (black-slash).
```
4) You can create problem testcase following this method:
 - Make new directory as `./problem/<problem_id>/`
 - Make output and input file as:
  - Input use `<number_testcase>.int`
  - Output use `<number_testcase>.out`
 - Make a description of problem as `desc.json` with these content:
```json
{
    "problem_title": "Problem_Name",
    "max_time": 1,
    "max_memory": 65536,
    "amount_testcase": 3
}
``` 
If you can't imagine what's you should to make them, You can use this picture as reference.

![Sample Folder Structure to add new problem list](https://ngixx.in.th/img/sample_go-grader.png)

_Note: max_time and max_memory are used seconds and kilobytes (kB)_

5) Let's start! *(with many bug!)*

# Cautions
This programs is in testing. It support for Debian.
Who's interest to use/contributed this script. You're welcome!