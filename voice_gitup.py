#!/usr/bin/env python3

import speech_recognition
import re
import subprocess
import pathlib
import os

recognizer = speech_recognition.Recognizer()
with speech_recognition.Microphone() as source:
    print("Say the magic words")
    audio = recognizer.listen(source)

words = recognizer.recognize_google(audio)

path = pathlib.Path().absolute()
matches = re.search("github (.*)", words)
if matches:
    os.chdir(path)
    print("Message is:", matches[1])
    subprocess.call(['gitup', '-b', "master", '-m', matches[1]])
else:
    print("Something went wrong")
