;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
; demo implementing a counter with closure ;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

( define make-counter                               ; make or reset counter by redefining the counter procedure
                                                    ; we quote the definition of counter to delay actual creation 
                                                    ; to when make-counter will be called
        '(define counter                            ; counter points to a procedure that increments counter
            (let ((count 0))                        ; Initialize count to 0 within the closure
                ( display "counter created/reset")
                (newline)   
                (lambda ()                          ; creates a procedure/closure that can be used to access and update count
                    (set! count (+ count 1))        ; Increment count by 1
                    count                           ; procedure will return the current count
                )                                   ; end of lambda definition
            )                                       ; end of let, returns the last argument, the procedure just constructed
        )                                           ; end of counter define, counter points to the procedure just constructed   
)                                                   ; end of make-counter definition     

; calling counter should fail, because none created yet
( display "calling counter should fail, because not yet created " ) ( newline)
( counter )

; first, create a counter
( newline) ( display "now, le'ts make a new counter") ( newline)
(make-counter) (newline)
( display (counter)) (newline)
( display (counter)) (newline)
( display (counter)) (newline)

; reset counter
( newline) ( display "resetting counter") ( newline)
( make-counter) ( newline)
( display (counter)) (newline)
( display (counter)) (newline)
( display (counter)) (newline)

( display "let's confirm that count is not accessible from outisde the closure ?") ( newline)
count 
( count )


