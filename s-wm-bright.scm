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

(define step 3)

(define brightnessctl-exec "/usr/bin/brightnessctl")

(define (up)
  (system*  brightnessctl-exec "set" (format #f "~a%+" step)))
(define (down)
  (system* brightnessctl-exec "set" (format #f "~a%-" step)))
(define (value options)
  (system* brightnessctl-exec "set" (cdr (assv 'value options))))

;; CLI PARSING
;; -----------------------------------------------------------------------

(define (cli-usage-banner)

  (display "bright [options]
  -u, --up              turn backlight up
  -d, --down            turn backlight down
  -v, --value           set backlight to value"))

(define cli-option-list '((up     (single-char #\u) (value #f))
                          (down   (single-char #\d) (value #f))
                          (value  (single-char #\v) (value #t))))

(define (cli-parser args)
  (let* ((option-spec cli-option-list)
         (options (getopt-long args option-spec)))
    (cli-option-run options)))

(define (cli-option-run options)
  (let ((option-wanted (lambda (option) (option-ref options option #f))))
    (cond ((option-wanted 'up)        (up))
          ((option-wanted 'down)      (down))
          ((option-wanted 'value)     (value options))
          (else                       (cli-usage-banner)))))

;; MAIN
;; -----------------------------------------------------------------------

(define (main args)
  (cli-parser args))
