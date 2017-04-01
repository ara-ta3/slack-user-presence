VIRTUALENV=$(shell which virtualenv)
env_dir=env
requirements=requirements.txt
newrelic=newrelic.ini


run:
	source $(env_dir)/bin/activate && NEW_RELIC_CONFIG_FILE=$(newrelic) python main.py

install: $(env_dir)
	source $(env_dir)/bin/activate && pip install -r $(requirements)

$(env_dir):
	$(MAKE) virtualenv

virtualenv:
	$(VIRTUALENV) -p python3 $(env_dir)
