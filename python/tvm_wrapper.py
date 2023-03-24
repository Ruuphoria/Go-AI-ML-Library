import model_loader
import nnvm
import os
import json


class TvmCompileParameter:
    def __init__(self, model_root_path, model_type, output_folder, shape_dict, data_dict="float32", target="llvm", opt_level=3):
        self.model_root_path = model_root_path
        self.model_type = model_type
        self.output_folder = output_folder
        self.target = target
        self.opt_level = opt_level
        self.data_dict = data_dict
        self.shape_dict = shape_dict


class TvmWrapper:
    def __init__(self, param):
        self.__param = param
        self.__model = None

        # フォルダが存在しない場合は作成する
        os.makedi