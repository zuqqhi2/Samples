#!/usr/bin/ruby

ARGF.each do |line|
	line.chomp!
	line.split.each do |word|
		puts "#{word}\t1"
	end
end
