#!/usr/bin/env ruby

require 'jimson'

class Shuffle
  extend Jimson::Handler 

  def Process(data) 
      puts "shuffle greeting"
      reply = {}
      if data.has_key?("greeting")
          s = data["greeting"]
          reply[:greeting] = s.split(//).sort_by { rand }.join('')
          reply[:shuffle] = "done"
      else
          reply[:shuffle] = "greeting not found"
      end
      return reply
  end
end

server = Jimson::Server.new(Shuffle.new, {:port => 8883})
server.start
