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
        os.makedirs(self.__param.model_root_path, exist_ok=True)
        os.makedirs(self.__param.output_folder, exist_ok=True)

    def setup(self):
        loader = model_loader.ModelLoaderFactory.get_loader(
            model_folder=self.__param.model_root_path,
            model_type=self.__param.model_type
        )
        self.__model = loader.load()

    def compile(self):
        # モデルからsymbol, parameterを取得
        # TODO : Darknetの対応が必要
        sym, params = nnvm.frontend.from_onnx(self.__model)

        # symbol, parameterからビルドを実施
        # TODO : Relayの対応
        with nnvm.compiler.build_config(opt_level=self.__param.opt_level):
            # build
            graph, lib, params = nnvm.compiler.build(
                graph=sym,
                target=self.__param.target,
                shape=self.__param.shape_dict,
                dtype=self.__param.data_dict,
 