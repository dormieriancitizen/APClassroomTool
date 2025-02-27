import json
import os

from colorama import Fore, Style, init

from bs4 import BeautifulSoup
from pathlib import Path

with open(Path("response.json"),"rb") as response:
    data = json.load(response)

    items = data["data"]["apiActivity"]["items"]

    i = 0

    for item in items:
        questions = item["questions"]
        for question in questions:
            # os.system('clear')
            i += 1

            if question["validation"]["scoring_type"] != "exactMatch":
                print(question["validation"]["scoring_type"])
                continue

            answers = {answer["value"]: answer["label"] for answer in question["options"]}

            answerCode = question["validation"]["valid_response"]["value"][0]
            answerHTML = answers[answerCode]
            
            answerSoup = BeautifulSoup(answerHTML,features="html.parser")
            answer = answerSoup.select(".choice_paragraph")[0].get_text()
            
            answerNumber = int(answerCode[-1])
            answerABC = "ABCDEFGHIJKLMNOP"[answerNumber-1]

            questionHTML = question["stimulus"]
            questionSoup = BeautifulSoup(questionHTML,features="html.parser")
            questionText = questionSoup.select(".stem_paragraph")[0].get_text()

            rationale = "None"
            if "custom_distractor_rationale_response_level" in question["metadata"]:
                rationaleList = question["metadata"]["custom_distractor_rationale_response_level"]
                rationaleHTML = rationaleList[answerNumber-1]
                rationaleSoup = BeautifulSoup(rationaleHTML,features="html.parser")
                rationale = rationaleSoup.text

            print(f"{Style.RESET_ALL}{Style.BRIGHT+Fore.RED}{i}. {Style.RESET_ALL}{questionText}")

            prettyInfo = [
                Fore.BLUE+Style.BRIGHT+answerABC,
                Fore.GREEN+answer,
                Style.DIM+rationale 
            ]

            sep = f"{Style.RESET_ALL+Style.DIM} - {Style.RESET_ALL}"
            print(sep + sep.join([line+"\n" for line in prettyInfo]))

            input("Press Enter to continue...")
