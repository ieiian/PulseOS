import CoreMotion
import Foundation

final class IOSStepReader {
    private let pedometer = CMPedometer()

    func isAvailable() -> Bool {
        CMPedometer.isStepCountingAvailable()
    }

    func previewSteps() -> Int {
        isAvailable() ? 5620 : 4200
    }

    var instance: CMPedometer {
        pedometer
    }
}

