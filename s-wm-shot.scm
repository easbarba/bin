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

;; VARIABLES
;; -----------------------------------------------------------------------

(define home (getenv "HOME"))
(define shot-folder (string-append home "/" "Pictures" "/" "screenshots" "/"))
(define shot-name (format #f "~a.png" (strftime "%A_%d_%B_%T" (localtime (current-time)))))

;; BRIGHTERS
;; -----------------------------------------------------------------------

;; GRIMSHOT
(define grimshot-exec "/usr/bin/grimshot")

;; ACTIONS
;; -----------------------------------------------------------------------

(define (full) (system*  grimshot-exec "save" "active" (string-append shot-folder shot-name)))
(define (partial) (system* grimshot-exec "save" "area" (string-append shot-folder shot-name)))

;; CLI PARSING
;; -----------------------------------------------------------------------

(define (cli-usage-banner)

  (display "shot [options]
  -f, --full              full screen shot
  -p, --partial           partial screen shot"))

(define cli-option-list '((full     (single-char #\f) (value #f))
                          (partial   (single-char #\p) (value #f))))

(define (cli-parser args)
  (let* ((option-spec cli-option-list)
         (options (getopt-long args option-spec)))
    (cli-option-run options)))

(define (cli-option-run options)
  (let ((option-wanted (lambda (option) (option-ref options option #f))))
    (cond ((option-wanted 'full)    (full))
          ((option-wanted 'partial) (partial))
          (else                     (cli-usage-banner)))))

;; MAIN
;; -----------------------------------------------------------------------

(define (main args)
  (cli-parser args))
