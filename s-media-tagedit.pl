#!/usr/bin/perl -w

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
