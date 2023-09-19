#!/usr/bin/guile \
--no-auto-compile -e main -s
!#

;; Licensed to the Apache Software Foundation (ASF) under one
;; or more contributor license agreements.  See the NOTICE file
;; distributed with this work for additional information
;; regarding copyright ownership.  The ASF licenses this file
;; to you under the Apache License, Version 2.0 (the
;; "License"); you may not use this file except in compliance
;; with the License.  You may obtain a copy of the License at
;;
;;  http://www.apache.org/licenses/LICENSE-2.0
;;
;; Unless required by applicable law or agreed to in writing,
;; software distributed under the License is distributed on an
;; "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
;; KIND, either express or implied.  See the License for the
;; specific language governing permissions and limitations
;; under the License.

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
