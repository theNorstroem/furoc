# furoc

```
furoc 
-I./node_modules/pathTto/spec/project
--plugin=furoc-gen-u33e
--u33e_out=
Sreference-search,\
Tform,\
Scollection-dropdown,\
:outputBaseDirectoryForU33e
```

```
furoc  //looks for .spectools in cwd

# Add a furoc section to your .spectools config
build:
  furoc:
    Input:
        - ./
    Commands:
      - OutputDir: dist/furoctest
        Plugin: furoc-gen-sample
        Args:
            - coldrun
```

--u33e_out= ==> furoc-gen-u33e binary
