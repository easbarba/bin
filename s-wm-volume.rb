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

require 'ostruct'
require 'optparse'

# Commands available
class Volume
  def initialize(options)
    @options = options
    @manager = Managers.new.first { |exec| File.executable?(exec.to_s) }
  end

  def final_command
    direction = @manager.output
    direction = @manager.input if @options[:direction] == :input

    case @options[:action]
    when :down
      direction.down
    when :up
      direction.up
    when :toggle
      direction.toggle
    end
  end

  def run
    system final_command
  end
end

# List of available sound managers
class Managers
  include Enumerable

  def initialize
    @known = [Pactl.new, Amixer.new, Mixer.new]
  end

  def each(&block)
    @known.each(&block)
  end
end

# pactl - Control a running PulseAudio sound server
class Pactl
  attr_reader :input, :output

  EXEC = '/usr/bin/pactl'
  STEP = 3

  def initialize
    @output = OpenStruct.new(output_data)
    @input = OpenStruct.new(input_data)
  end

  def to_s
    <<~INFO
      exec: #{EXEC}
      sink: #{sink}

      -- spec --

      #{`pactl info`}
    INFO
  end

  private

  def output_data
    output_id = '@DEFAULT_SINK@'
    { toggle: "#{EXEC} set-sink-mute #{output_id} toggle",
      up: "#{EXEC} set-sink-volume #{output_id} +#{STEP}%",
      down: "#{EXEC} set-sink-volume #{output_id} -#{STEP}%" }
  end

  def input_data
    input_id = '@DEFAULT_SOURCE@'
    { toggle: "#{EXEC} set-source-mute #{input_id} toggle",
      up: "#{EXEC} set-source-volume #{input_id} +#{STEP}%",
      down: "#{EXEC} set-source-volume #{input_id} -#{STEP}%" }
  end
end

# bsd mixer
class Mixer
  EXEC = '/sbin/mixer'

  def mixer
    {
      name: 'mixer',
      toggle: '',
      updown: "mixer vol #{states[state]}#{STEP}"
    }
  end

  def toggle
    [mixer[:name], mixer[:toggle]].join(' ')
  end

  def updown
    [mixer[:name], mixer[:updown]].join(' ')
  end
end

# Command-line ALSA mixer
class Amixer
  EXEC = '/usr/bin/amixer'

  def amixer
    {
      name: 'amixer',
      toggle: '-q sset Master toggle',
      updown: "set Master #{STEP}%#{states[state]}"
    }
  end

  def toggle
    [amixer[:name], amixer[:toggle]].join(' ')
  end

  def updown
    [amixer[:name], amixer[:updown]].join(' ')
  end
end

# hail to the new king
class Pipewire
  def info
    {
      name: 'pw-cli',
      args: "s #{pipe_id} Props",
      toggle: '{ mute: false, channelVolumes: [ 1.5, 1.5 ] }',
      updown: '{ mute: false, channelVolumes: [ 1.5, 1.5 ] }'
    }
  end

  def id; end

  def toggle
    [info[:name], info[:toggle]].join(' ')
  end

  def updown
    [info[:name], info[:updown]].join(' ')
  end
end

options = {}
oparser = OptionParser.new do |opts|
  opts.banner = 'Usage: volume [options]'

  opts.on('-u', '--up', 'Increase volume') do
    options[:action] = :up
  end

  opts.on('-d', '--down', 'Decrease volume') do
    options[:action] = :down
  end

  opts.on('-t', '--toggle', 'Toggle volume') do
    options[:action] = :toggle
  end

  opts.on('-i', '--input', 'Manage input settings') do
    options[:direction] = :input
  end

  opts.on('-o', '--output', 'Manage output settings') do
    options[:direction] = :output
  end
end

oparser.parse! ['--help'] if ARGV.empty?
oparser.parse!

Volume.new(options).run
