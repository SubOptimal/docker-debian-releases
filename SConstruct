
if False:
    Environment = None

env=Environment()
env.Command("_apt-mirrors.json", [], "./apt-mirrors-info.go $TARGET")
env.Command("_travis.yml", ["_apt-mirrors.json", "travis.yml.tmpl"],
            "./json-tmpl-render.go $SOURCES $TARGET")
