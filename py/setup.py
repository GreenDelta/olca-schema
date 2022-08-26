from setuptools import setup, find_packages
from os import path

here = path.abspath(path.dirname(__file__))

# Get the long description from the README file
with open(path.join(here, 'README.md'), encoding='utf-8') as f:
    long_description = f.read()

setup(
    name='olca_schema',
    version='0.0.3',
    description='A package for reading and writing data sets in the openLCA schema format.',
    long_description=long_description,
    long_description_content_type='text/markdown',
    url='https://github.com/GreenDelta/olca-schema',
    packages=find_packages(exclude=[
        "tests", "*.tests", "*.tests.*", "tests.*"]),
    keywords=['openLCA', 'life cycle assessment', 'LCA'],
    license="CC0",
    classifiers=[
        "Development Status :: 2 - Pre-Alpha",
        "Environment :: Console",
        "Intended Audience :: Science/Research",
        "License :: CC0 1.0 Universal (CC0 1.0) Public Domain Dedication",
        "Programming Language :: Python :: 3.8",
        'Programming Language :: Python :: 3.9',
        'Programming Language :: Python :: 3.10',
        "Topic :: Utilities",
    ]
)
