#!/usr/bin/guile \
--no-auto-compile -e main -s
!#

;; Bin is free software: you can redistribute it and/or modify
;; it under the terms of the GNU General Public License as published by
;; the Free Software Foundation, either version 3 of the License, or
;; (at your option) any later version.
;;
;; Bin is distributed in the hope that it will be useful,
;; but WITHOUT ANY WARRANTY; without even the implied warranty of
;; MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
;; GNU General Public License for more details.
;;
;; You should have received a copy of the GNU General Public License
;; along with Bin. If not, see <https://www.gnu.org/licenses/>.

(import (rnrs base)
        (rnrs io ports)
        (rnrs io simple)
        (ice-9 getopt-long)
        (ice-9 ftw)
        (ice-9 format))

;; ACTIONS
;; -----------------------------------------------------------------------

(define step 1)

(define pactl-exec "/usr/bin/pactl")
(define pactl-id "@DEFAULT_SINK@")

(define (up) (system*  pactl-exec "set-sink-volume" pactl-id (format #f "+~a%" step)))
(define (down) (system* pactl-exec "set-sink-volume" pactl-id (format #f "-~a%" step)))
(define (toggle) (system* pactl-exec "set-sink-mute" pactl-id "toggle"))

;; CLI PARSING
;; -----------------------------------------------------------------------

(define (cli-usage-banner)

  (display "volume [options]
  -u, --up              increase volume
  -d, --down            decrease volume
  -t, --toggle          toggle volume"))

(define cli-option-list '((up       (single-char #\u) (value #f))
                          (down     (single-char #\d) (value #f))
                          (toggle   (single-char #\t) (value #f))))

(define (cli-parser args)
  (let* ((option-spec cli-option-list)
         (options (getopt-long args option-spec)))
    (cli-option-run options)))

(define (cli-option-run options)
  (let ((option-wanted (lambda (option) (option-ref options option #f))))
    (cond ((option-wanted 'up)        (up))
          ((option-wanted 'down)      (down))
          ((option-wanted 'toggle)    (toggle))
          (else                       (cli-usage-banner)))))

;; MAIN
;; -----------------------------------------------------------------------

(define (main args)
  (cli-parser args))
