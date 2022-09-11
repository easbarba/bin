#!/usr/bin/env ruby
# frozen_string_literal: true

require 'pathname'
require 'fileutils'
require 'optparse'

# TODO: add --verbose option

# Take a shot of the marvelous screen
class Screenshot
  def self.which?(executable)
    ENV['PATH'].split(File::PATH_SEPARATOR).any? do |directory|
      File.executable?(File.join(directory, executable.to_s))
    end
  end

  attr_reader :shotters

  def initialize(shotters = Shotters.new.all)
    @shotters = shotters
  end

  def full
    [shotter.exec, shotter.common, shotter.full, shotter.name].join ' '
  end

  def partial
    [shotter.exec, shotter.common, shotter.partial, shotter.name].join ' '
  end

  def info
    p <<~EOL # double quote document
      Shotters available: #{shotters.keys.join(', ')}
      location: #{shotter.name}
    EOL
  end

  private

  # list of all shotters found
  def shotters_available
    shotters.find_all do |shotter|
      x = shotter.last
      x.exec if Screenshot.which? x.exec
    end
  end

  # select default shotter
  def shotter
    shotters_available.first.last
  end
end

# All available shotters
class Shotters
  require 'ostruct'

  PREFERRED_FORMAT = 'png'
  CURRENT_TIME = Time.new.strftime '%d-%m-%y at %I-%M'

  def all
    {
      scrot: OpenStruct.new(exec: 'scrot',
                            full: '--focused --silent', partial: '--select --silent', common: '--border --quality=100',
                            name: screenshot_location),
      main: OpenStruct.new(exec: 'maim',
                           full: '', partial: '--select',
                           name: screenshot_location),
      flameshot: OpenStruct.new(exec: 'flameshot',
                                full: '', partial: '--select',
                                name: screenshot_location),
      grim: OpenStruct.new(exec: 'grim',
                           full: '', partial: '-g "$(slurp)"', common: '-t png -l 9',
                           name: screenshot_location)
    }
  end

  private

  def folder
    f = Pathname.new(Dir.home).join('Pictures', 'screenshots')
    FileUtils.mkdir_p f unless f.exist?

    f
  end

  def screenshot_location
    "'#{folder.join("Screenshot from #{CURRENT_TIME}.#{PREFERRED_FORMAT}")}'"
  end
end

options = OptionParser.new do |opts|
  shotter = Screenshot.new
  opts.banner = 'Usage: shot [options]'

  opts.on('-f', '--full', 'full screen shot') do
    system shotter.full
  end

  opts.on('-p', '--partial', 'partial screen shot') do
    system shotter.partial
  end

  opts.on('-i', '--info', 'shotters information') do
    puts shotter.info
  end
end

options.parse! ['--help'] if ARGV.empty?
options.parse!
