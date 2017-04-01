from newrelic import agent

if __name__ == '__main__':
    application = agent.register_application()
    print(application)
    res = agent.record_custom_metric('Custom/Value', 100, application)
    print(res)
