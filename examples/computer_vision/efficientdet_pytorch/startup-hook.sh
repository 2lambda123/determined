pip install timm==0.9
# pycocotools 2.0.5, a dependency of efficientdet-pytorch,
# would not install without cython
pip install pycocotools==2.0.6
git clone https://github.com/rwightman/efficientdet-pytorch.git
cd efficientdet-pytorch
git checkout d43c9e34cd62d22b4205831bb735f6dd83b8e881
python setup.py install
cd ..
