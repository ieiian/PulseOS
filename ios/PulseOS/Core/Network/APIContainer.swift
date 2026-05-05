import Foundation

final class APIContainer: ObservableObject {
    let client: APIClient
    let home: HomeService
    let diet: DietService
    let activity: ActivityService
    let meditation: MeditationService
    let sleep: SleepService
    let user: UserService

    static let shared = APIContainer()

    init(client: APIClient = APIClient()) {
        self.client = client
        self.home = HomeService(client: client)
        self.diet = DietService(client: client)
        self.activity = ActivityService(client: client)
        self.meditation = MeditationService(client: client)
        self.sleep = SleepService(client: client)
        self.user = UserService(client: client)
    }
}
