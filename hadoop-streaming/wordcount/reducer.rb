#!/usr/bin/ruby

count = Hash.new {|h,k| h[k] = 0}
ARGF.each do |line|
	line.chomp!
	key, value = line.split(/\t/)
	count[key] += 1
end

count.each do |k,v|
	puts "#{k}\t#{v}"
end
