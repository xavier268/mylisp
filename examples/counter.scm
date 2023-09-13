;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;
; demo implementing a counter with closure ;
;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;;

( define make-counter                       ; make or reset counter by redefining the counter procedure
                                            ; we quote the definition of counter to delay actual creation
        '(define counter                    ; counter will increment the created counter
            (let ((count 0))                ; Initialize count to 0 within the closure
                ( display "counter created/reset")
                (newline)   
                (lambda ()                  ; Return a closure that can be used to access and update count
                (set! count (+ count 1))    ; Increment count by 1
                count                       ; Return the current count
                )                           ; end of lambda definition
            )                               ; end of let
        )                                   ; end of define    
)                                           ; end of reset-counter          

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

(quit)

