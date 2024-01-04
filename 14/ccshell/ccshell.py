import subprocess
import signal
import os

history = []

while True:
    prev = None
    signal.signal(signal.SIGINT, lambda sig, _: 0)
    cli = input("ccshell> ")
    cmds = cli.split("|")
    cmds = [cmd.strip() for cmd in cmds]
    for idx in range(len(cmds)):
        cl = cmds[idx].split(" ")
        cl = [c.strip() for c in cl]
        match cl:
            case ["exit"]:
                exit()
            case [""]:
                pass
            case ["history"]:
                for h in history:
                    print(h)
            case ["cd", args]:
                try:
                    os.chdir("".join(args))
                except Exception as e:
                    print('No such file or directory ' + "".join(args) + '\n', end="")
            case ["pwd"]:
                print(os.getcwd())
            case cmd:
                try:
                    if prev is None and idx == len(cmds) - 1:
                        p = subprocess.Popen(cmd)
                        p.wait()
                    elif prev and idx == len(cmds) - 1:
                        p = subprocess.Popen(cmd, stdin=prev.stdout, text=True)
                        p.wait()
                        output = p.communicate()
                    elif prev is not None:
                        p = subprocess.Popen(cmd, stdin=prev.stdout, stdout=subprocess.PIPE, text=True)
                    else:
                        p = subprocess.Popen(cmd, stdout=subprocess.PIPE, text=True)
                    prev = p
                    signal.signal(signal.SIGINT, lambda sig, _: p.terminate())
                except Exception as e:
                    print('No such file or directory (os error 2)\n', end="")
    history.append(cli)
