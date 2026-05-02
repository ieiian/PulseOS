import Foundation

final class SleepMonitorController: ObservableObject {
    @Published private(set) var isRecording = false

    func start() {
        isRecording = true
    }

    func stop() {
        isRecording = false
    }
}

