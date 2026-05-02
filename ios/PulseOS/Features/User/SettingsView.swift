import SwiftUI

struct SettingsView: View {
    var body: some View {
        VStack(alignment: .leading, spacing: 12) {
            Text("设置")
                .font(.title2.weight(.semibold))
            Text("通知设置")
            Text("权限管理")
            Text("设备连接")
            Spacer()
        }
        .padding(16)
        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .topLeading)
        .background(PulseTheme.background)
    }
}

