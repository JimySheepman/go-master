import pandas as pd
import os

input_file_path = "file_path"

excel_file_list = os.listdir(input_file_path)

excel_file_list

df = pd.DataFrame()

for excel_files in excel_file_list:
    if excel_files.endswith(".xlsx"):
        df1 = pd.read_excel(input_file_path+excel_files)
        df = df.append(df1)

df.to_excel(input_file_path+"PDR.xlsx")
