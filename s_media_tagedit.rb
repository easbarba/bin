#!/usr/bin/env ruby
# frozen_string_literal: true
#
# Bin is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Bin is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Bin. If not, see <https://www.gnu.org/licenses/>.
#

# TODO: permited media files extension
# TODO: walk through all sub-directories if required

# Requirements
# gem install taglib-ruby

require 'taglib'
require 'pathname'
require 'optparse'

TITLE = lambda do |info, file|
  info = Pathname.new(file).basename('.*').to_path if info.empty?
  TagLib::FileRef.open(file) do |t|
    t.tag.title = info
    t.save
  end
end

ARTIST = lambda do |info, file|
  TagLib::FileRef.open(file) do |t|
    t.tag.artist = info
    t.save
  end
end

ALBUM = lambda do |info, file|
  TagLib::FileRef.open(file) do |t|
    t.tag.album = info
    t.save
  end
end

def run(filepath, info, func)
  raise 'FILE NOT FOUND' unless filepath.exist?

  if filepath.directory?
    filepath.children.each { |file| func.call(info, file.to_path) }
  else
    func.call(info, filepath.to_path)
  end
end

option_parser = OptionParser.new do |opts|
  opts.banner = 'Usage: tagedit [options]'
  filepath = String.new

  opts.on('-fFILE', '--file=FILE', 'file to edit') do |file|
    filepath = Pathname.new(file)
  end

  opts.on('-tTITLE', '--tittle=TITLE', 'edit title tag') do |title|
    run(filepath, title, TITLE)
  end

  opts.on('-aNAME', '--album=NAME', 'edit album tag') do |info|
    run(filepath, info, ALBUM)
  end

  opts.on('-sNAME', '--artist=NAME', 'edit artist tag') do |info|
    run(filepath, info, ARTIST)
  end
end

option_parser.parse! ['--help'] if ARGV.empty?
option_parser.parse!
