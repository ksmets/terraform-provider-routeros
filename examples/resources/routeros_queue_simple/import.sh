#The ID can be found via API or the terminal
#The command for the terminal is -> :put [/queue/simple get [print show-ids]]
terraform import routeros_queue_simple.test *3
#Or you can import a resource using one of its attributes
terraform import routeros_queue_simple.test "name=server"