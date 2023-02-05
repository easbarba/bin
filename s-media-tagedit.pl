#!/usr/bin/perl -w

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


use v5.32;
use utf8;
use warnings;
use strict;
use Getopt::Long 'HelpMessage';
use Time::Piece;
use File::Find;
use Audio::TagLib;

my $f      = Audio::TagLib::FileRef->new("Latex Solar Beef.mp3");
my $artist = $f->tag()->artist();
print $artist->toCString(), "\n"; # got "Frank Zappa"

$f->tag()->setAlbum(Audio::TagLib::String->new("Fillmore East"));
$f->save();

my $g      = Audio::TagLib::FileRef->new("Free City Rhymes.ogg");
my $album  = $g->tag()->album();
print $album->toCString(), "\n";  # got "NYC Ghosts & Flowers"

$g->tag()->setTrack(1);
$g->save();
