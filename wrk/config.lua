
math.randomseed(os.time())
math.random(); math.random(); math.random()

token = 'OTIwNjgwYjUtNmY5OC00NmYyLTk3NDQtNzEzNjY3MzkzM2Ji'

request = function()
  path = "/api/admin/1/polls/1/constructor"
  return wrk.format(
          'GET',
          path,
          {
            ['Host'] = 'localhost',
            ["Content-Type"] = "application/json",
            ["Cookie"] = "OPROSSO_SESSION_ID=" .. token .. ""
          })
end

response = function(status, headers, body)
  for key, value in pairs(headers) do
    if key == "Location" then
      io.write("Location header found!\n")
      io.write(key)
      io.write(":")
      io.write(value)
      io.write("\n")
      io.write("---\n")
      break
    end
  end
end
