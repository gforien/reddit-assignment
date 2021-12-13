BeforeAll {
    # Start Go server and store process object
    $proc = Start-Process '.\reddit-assignment.exe' -PassThru
}

Describe 'Test reddit-assignment.go' {
# NB: Eventhough Pester is a framework for unit-testing Powershell code,
#     it is used here for integration testing.

    It "should GET /ok" {
        $res = iwr 'http://localhost:5000/ok' -method GET
        $res.StatusCode | Should -Be 200
        $res.Content | Should -Be "gin OK"
    }
    It "should GET /count" {
        $res = iwr 'http://localhost:5000/count' -method GET
        $res.StatusCode | Should -Be 200
        $res.Content | Should -Be 0
    }
    It "should POST /inc" {
        $res = iwr 'http://localhost:5000/inc' -method POST
        $res.StatusCode | Should -Be 200
        $res.Content | Should -Be 1
    }
    It "should POST /inc" {
        $res = iwr 'http://localhost:5000/inc' -method POST
        $res.StatusCode | Should -Be 200
        $res.Content | Should -Be 2
    }
    It "should POST /inc" {
        $res = iwr 'http://localhost:5000/inc' -method POST
        $res.StatusCode | Should -Be 200
        $res.Content | Should -Be 3
    }
    It "should POST /dec" {
        $res = iwr 'http://localhost:5000/dec' -method POST
        $res.StatusCode | Should -Be 200
        $res.Content | Should -Be 2
    }
    It "should GET /count" {
        $res = iwr 'http://localhost:5000/count' -method GET
        $res.StatusCode | Should -Be 200
        $res.Content | Should -Be 2
    }
}

AfterAll {
    # Teardown the process started in BeforeAll block
    $proc | Stop-Process
}
